import React from 'react';
import { Router, Route, browserHistory } from 'react-router';
import { Provider } from 'react-redux';
import { MainPage, LoginPage, ProfilePage, ListPage } from './pages';

export default class App extends React.Component {
  render() {
    return (
      <Provider store={this.props.store}>
        <Router history={browserHistory}>
          <Route path="/" component={MainPage} store={this.props.store} />
          <Route path="/login" component={LoginPage} store={this.props.store} />
          <Route path="/profile" component={ProfilePage} store={this.props.store} />
          <Route path="/list" component={ListPage} store={this.props.store} />
        </Router>
      </Provider>
    );
  }
}
