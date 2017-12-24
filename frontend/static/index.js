var app = angular.module("app", []);
app.controller('lucas', ['$scope', '$http', function ($scope, $http) {
    $http({
        method: 'GET',
        url: "/store"
    }).success(function (data) {
        $scope.nodes = data;
    });

    $scope.update = function (newValue) {
        console.log(newValue);
    }
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
            "<p class='item-key' ng-click='hide=!hide' ng-class=\"{true: 'has-key', false: 'not-has-key'}[item.KV!=undefined]\">{{ item.RootPath}}</p>" +
            "<p ng-if='item.KV != undefined' class='item-value'><button item-detail class='btn btn-primary btn-xs' kv='item.KV'>view</button></p>" +
            "<div ng-show='hide==false' ng-if='hasNodes(item.Nodes)' style='margin-left:16px'><tree-view nodes='item.Nodes'></tree-view></div>" +
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


app.directive("itemDetail", ['$rootScope', function ($rootScope) {
    function link(scope, element, attrs) {
        element.on("click", function () {
            $rootScope.json = scope.kv;
            $rootScope.$apply()
        })
    }

    return {
        restrict: 'AE',
        scope: {
            "kv": "="
        },
        link: link
    }
}]);

app.directive("updateItem", ['$http', function ($http) {
    function link(scope, element, attrs) {
        element.on("click", function () {
            console.log(scope.kv);
            $http(
                {
                    method: "POST",
                    url: "/store",
                    data: {
                        "key": scope.kv.key,
                        "value": scope.kv.value
                    },
                    headers: {'Content-Type': 'application/x-www-form-urlencoded'},
                    transformRequest: function (obj) {
                        var str = [];
                        for (var p in obj)
                            str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
                        return str.join("&");
                    },
                }
            ).success(function () {
                alert("change value successfully")
            }).error(function () {
                alert("change value failed")
            })

        })
    }

    return {
        restrict: 'AE',
        scope: {
            "kv": "="
        },
        link: link
    }
}]);