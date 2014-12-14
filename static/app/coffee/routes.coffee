define [
  "angular"
  "app"
], (angular, app) ->
  app.config [
    "$routeProvider"
    ($routeProvider) ->
      $routeProvider.when "/view1",
            templateUrl: "app/partials/partial1.html"
            controller: "MyCtrl1"

      $routeProvider.when "/view2",
        templateUrl: "app/partials/partial2.html"
        controller: "MyCtrl2"

      $routeProvider.otherwise redirectTo: "/view1"
  ]
