define [
  "angular"
  "services"
], (angular, services) ->

  # Directives
  angular.module("myApp.directives", ["myApp.services"]).directive "appVersion", [
    "version"
    (version) ->
      return (scope, elm, attrs) ->
        elm.text version
        return
  ]
  return
