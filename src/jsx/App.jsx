import React from 'react';
import ReactDOM from 'react-dom';
import { Router, IndexRoute, Route, browserHistory } from 'react-router';
import { MainPage, LoginPage } from 'containers';

ReactDOM.render(
    <Router history={browserHistory}>
        <Route path='/' component={ MainPage }></Route>
        <Route path='/login' component={ LoginPage }></Route>
    </Router>,
    document.getElementById('app')
)
