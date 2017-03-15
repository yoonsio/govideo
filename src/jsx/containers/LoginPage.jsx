import React from 'react';
import { Content } from 'components';

class LoginForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            username: '',
            password: ''
        };
    }

    login = (e) => {
        // TODO: some validation
        e.preventDefault();
        console.log('triggering login with ' + this.state.username + ', ' + this.state.password)
        var form = new FormData();
        form.append('username', this.state.username);
        form.append('password', this.state.password);
        var request = new Request('/login', {
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
        fetch(request).then(function(response) {
            // perform setState here
            console.log('status: ' + response.status);
            return response.json();
        }).then(function(j) {
            console.log('json: ' + j);
        }).catch(function(err) {
            // error
            console.log(err);
        });
    }

    onChange = (e) => {
        this.setState({[e.target.name]: e.target.value});
    }

    render() {
        return (
            <form id="loginForm">
                <h2>Login</h2>
                <input id="login_username" name="username" type="email" value={this.props.username} onChange={this.onChange} placeholder="Email" required />
                <input id="login_password" name="password" type="password" value={this.props.password} onChange={this.onChange} placeholder="Password" required />
                <button id="signin-btn" className="btn btn-primary btn-sm" onClick={this.login}>SIGN IN</button>
            </form>
        )
    }

}

export default class LoginPage extends React.Component {
    render() {
        return (
            <Content>
                <LoginForm />
            </Content>
        )
    }
}
