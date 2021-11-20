import React from "react";
import { Switch, Route } from "react-router-dom";
import home from "../components/home"
import longin from "../components/login"
import article from "../components/article"
import banner from "../components/banner"
const routes = [
  {
    path: '/',
    component: home,
    exact: true,
    routes: []
  },
  {
    path: '/login',
    component: longin,
    exact: true,
    routes: []
  },
  {
    path: '/article',
    component: article,
    exact: true,
    routes: []
  },
  {
    path: '/banner',
    component: banner,
    exact: true,
    routes: []
  }
]

export default RouteConfig
function RouteConfig() {
  return (
    <Switch>
      {routes.map((route) => (
        <RouteWithSubRoutes key={route.path} {...route} />
      ))}
    </Switch>
  );
}

function RouteWithSubRoutes(route) {
  return (
    <Route
      path={route.path}
      render={props => (
        <route.component {...props} routes={route.routes} />
      )}
    />
  )
}
