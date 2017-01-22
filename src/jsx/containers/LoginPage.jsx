import React from 'react';
import { Content } from 'components';

const LoginForm = React.createClass({
    render() {
        return (
            <form id="loginForm">
                <h2>Login</h2>
                <input id="login_username" name="username" type="email" placeholder="Email" required />
                <input id="login_password" name="password" type="password" placeholder="Password" required />
                <button id="signin-btn" class="btn btn-primary btn-sm">SIGN IN</button>
            </form>
        )
    }
});

export default class LoginPage extends React.Component {
    render() {
        return (
            <Content>
                <LoginForm />
            </Content>
        )
    }
}
