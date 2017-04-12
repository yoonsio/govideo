import React from 'react';
import { Router, Route, IndexRoute, browserHistory } from 'react-router';
import { Provider } from 'react-redux';
import { Content } from './components';
import { MainPage, LoginPage, ProfilePage, ListPage, ViewPage } from './pages';

export default class App extends React.Component {
  render() {
    return (
      <Provider store={this.props.store}>
        <Router history={browserHistory}>
          <Route path="/" component={Content} store={this.props.store} >
            <IndexRoute component={MainPage} store={this.props.store} />
            <Route path="profile" component={ProfilePage} store={this.props.store} />
            <Route path="login" component={LoginPage} store={this.props.store} />
            <Route path="media" store={this.props.store} >
              <IndexRoute component={ListPage} store={this.props.store} />
              <Route path=":path" component={ViewPage} store={this.props.store} />
            </Route>
          </Route>
          {/*<Route path="*" component={NotFound} store={this.props.store} />*/}
        </Router>
      </Provider>
    );
  }
}
