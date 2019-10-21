package frontend

const HOME_PAGE = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
	<script src="https://cdnjs.cloudflare.com/ajax/libs/angular.js/1.2.20/angular.min.js"></script>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.2.0/css/bootstrap.min.css"/>
	<script src="https://code.jquery.com/jquery-3.2.1.min.js"></script>
	<script src="http://ringtail-lucas.oss-cn-beijing.aliyuncs.com/index.js"></script>
	<link rel="stylesheet" href="http://ringtail-lucas.oss-cn-beijing.aliyuncs.com/index.css"/>
    <title>Lucas</title>
</head>
<body ng-app="app">
	<div ng-controller="lucas" class="row main-container">
		<div class="col-md-12 container-title">Lucas - etcd v3 key value browser</div>
		<div class="col-md-8">
			<tree-view nodes="nodes.Nodes"></tree-view>
		</div>
		<div class="col-md-4 form" style="min-height:500px" ng-if="json">
			<textarea class="data-viewer form-control" ng-model="json.value"></textarea>
			<button update-item kv="json" class="btn btn-primary pull-right" style="margin-right:5%;margin-top:8px;">change</button>
		<div>
	</div>
</body>
</html>
`
