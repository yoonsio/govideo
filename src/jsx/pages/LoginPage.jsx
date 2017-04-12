import React from 'react';
import { connect } from 'react-redux';
import { browserHistory } from 'react-router';
import { Content, LoginForm } from '../components';

class LoginPage extends React.Component {

  componentWillMount() {
    const { user } = this.props;
    if (user != null) {
      if (Object.keys(user).length !== 0 && user.constructor === Object) {
        browserHistory.push('/');
      }
    }
  }

  render() {
    return (
        <LoginForm />
    );
  }
}

const mapStateToProps = state => ({
  user: state.user,
});

const ConnectedLoginPage = connect(mapStateToProps)(LoginPage);
export default ConnectedLoginPage;
