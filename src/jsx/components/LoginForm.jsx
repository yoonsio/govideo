import React from 'react';
import { connect } from 'react-redux';
import { loginUser } from '../helpers/userReq';

class LoginForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      username: '',
      password: '',
    };
  }

  onChange = (e) => {
    this.setState({ [e.target.name]: e.target.value });
  }

  login = (e) => {
    // TODO: some validation
    e.preventDefault();
    loginUser(this.props, this.state.username, this.state.password);
  }

  render() {
    return (
      <form id="loginForm">
        <h2>Login</h2>
        <input id="login_username" name="username" type="email" value={this.props.username} onChange={this.onChange} placeholder="Email" required />
        <input id="login_password" name="password" type="password" value={this.props.password} onChange={this.onChange} placeholder="Password" required />
        <button id="signin-btn" className="btn btn-primary btn-sm" onClick={this.login}>SIGN IN</button>
      </form>
    );
  }
}

const mapStateToProps = state => ({
  user: state.user,
});

const ConnectedLoginForm = connect(mapStateToProps)(LoginForm);
export default ConnectedLoginForm;
