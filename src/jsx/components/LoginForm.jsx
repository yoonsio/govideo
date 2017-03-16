import React from 'react';

export default class LoginForm extends React.Component {
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
    const form = new FormData();
    form.append('username', this.state.username);
    form.append('password', this.state.password);
    const request = new Request('/login', {
      method: 'post',
      body: form,
      credentials: 'same-origin',
      mode: 'cors',
      redirect: 'follow',
      cache: 'no-cache',
      headers: new Headers({
        'Content-Type': 'text/plain'
      })
    });
    fetch(request).then((response) => {
      // perform setState here
      console.log('status: ' + response.status);
      return response.json();
    }).then((j) => {
      console.log('json: ' + j);
    }).catch((err) => {
      // error
      console.log(err);
    });
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