import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Route, browserHistory } from 'react-router';
import { MainPage, LoginPage } from 'containers';

ReactDOM.render(
  <Router history={browserHistory}>
    <Route path="/" component={MainPage} />
    <Route path="/login" component={LoginPage} />
  </Router>,
  document.getElementById('app'),
);
