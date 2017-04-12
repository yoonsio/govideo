import React from 'react';
import { Link } from 'react-router';
import { connect } from 'react-redux';
import { logout } from '../helpers/userReq';

class Navbar extends React.Component {

  handleLogout = () => {
    logout(this.props);
  }

  renderUserComponent() {
    const { user } = this.props;
    console.log(user);
    if (user === null) {
      return <Link className="navbar-right-menu" to="/profile"><i className="fa fa-fw fa-refresh fa-spin" /> Loading</Link>;
    } else if (Object.keys(user).length === 0 && user.constructor === Object) {
      return (
        <div>
          <Link className="navbar-right-menu" to="/login"><i className="fa fa-fw fa-user" /> Login</Link>
        </div>
      );
    }
    return (
      <div>
        <Link className="navbar-right-menu" to="/profile">{user.first_name} {user.last_name} ({user.email})</Link>
        <Link className="navbar-right-menu" onClick={this.handleLogout}><i className="fa fa-fw fa-sign-out" /></Link>
      </div>
    );
  }

  render() {
    return (
      <nav className="navbar navbar-default">
        <div className="container-fluid">
          <div className="navbar-header">
            <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#main-navbar-collapse" aria-expanded="false">
              <span className="sr-only">Toggle navigation</span>
              <span className="icon-bar" />
              <span className="icon-bar" />
              <span className="icon-bar" />
            </button>
            <a className="navbar-brand" href="/"><i className="fa fa-fw fa-video-camera" /> GoVideo</a>
          </div>
          <div className="collapse navbar-collapse" id="main-navbar-collapse">
            <ul className="nav navbar-nav">
              <li><Link to="/media"><i className="fa fa-fw fa-play" /> My Library</Link></li>
              <li><Link to="/"><i className="fa fa-fw fa-gear" /> Setting</Link></li>
            </ul>
            <ul className="nav navbar-nav navbar-right">
              <li>
                <div>{this.renderUserComponent()}</div>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    );
  }
}

const mapStateToProps = state => ({
  user: state.user,
});

const ConnectedNavbar = connect(mapStateToProps)(Navbar);
export default ConnectedNavbar;
