import React from 'react';
import { connect } from 'react-redux';
import { Content } from '../components';

class ProfilePage extends React.Component {

  detail = () => {
    if (this.props.user) {
      return (
        <ul>
          <h1>
            {this.props.user.first_name} 
            &nbsp;
            {this.props.user.last_name}
            &nbsp;
            ({this.props.user.email})
          </h1>
        </ul>
      );
    }
    // TODO: redirect to login page
    // browserHistory.push('/login');
    return <h3>Please Log-in</h3>;
  }

  render() {
    return (
      <div>
        {this.detail()}
      </div>
    );
  }
}

const mapStateToProps = state => ({
  user: state.user,
});

const ConnectedProfilePage = connect(mapStateToProps)(ProfilePage);
export default ConnectedProfilePage;
