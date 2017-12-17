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
	<script src="https://code.jquery.com/jquery-3.2.1.min.js"></script>
	<script src="http://localhost:8081/static/index.js"></script>
	<link type="text/css" href="http://localhost:8081/static/index.css"/>
    <title>Document</title>
</head>
<body ng-app="app">
	<div ng-controller="lucas">
		<tree-view nodes="nodes.Nodes"></tree-view>
	</div>
</body>
</html>
`
