import React from 'react';
import { Router, Route, browserHistory } from 'react-router';
import { Provider } from 'react-redux';
import { MainPage, LoginPage } from './pages';

export default class App extends React.Component {
  render() {
    return (
      <Provider store={this.props.store}>
        <Router history={browserHistory}>
          <Route path="/" component={MainPage} store={this.props.store} />
          <Route path="/login" component={LoginPage} store={this.props.store} />
          <Route path="/profile" component={LoginPage} store={this.props.store} />
        </Router>
      </Provider>
    );
  }
}
