var app = angular.module("app", []);
app.controller('lucas', ['$scope', '$http', function ($scope, $http) {
    $http({
        method: 'GET',
        url: "/store"
    }).success(function (data) {
        $scope.nodes = data;
    })
}]);

app.directive("treeView", ["$compile", function ($compile) {
    function link(scope, element, attrs) {
        scope.hasNodes = function (nodes) {
            if (nodes != undefined && Object.keys(nodes).length > 0) {
                return true;
            } else {
                return false;
            }
        };
        element.append("<div ng-repeat='item in nodes'>" +
            "{{ item.RootPath}}" +
            "<span ng-if='item.KV != undefine'>: {{item.KV.value}}</span>" +
            "<div ng-if='hasNodes(item.Nodes)' style='margin-left:6px'><tree-view nodes='item.Nodes'></tree-view></div>" +
            "</div>"
        );
        $compile(element.contents())(scope);
    }

    return {
        restrict: 'AE',
        scope: {
            "nodes": "="
        },
        link: link
    };
}]);
