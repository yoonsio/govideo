import React from 'react';
import { connect } from 'react-redux';
import { Content } from '../components';

class ViewPage extends React.Component {

    constructor(props) {
        super(props);
        this.stae = {
            path: null,
        }
    }

    render() {
        var fixedWidth = {
            width: "300px"
        };
        return (
            <div>
                <p>{`/media/${this.props.params.path}/data`}</p>
            </div>
        )
    }
}

const mapStateToProps = state => ({
  user: state.user,
});

const ConnectedViewPage = connect(mapStateToProps)(ViewPage);
export default ConnectedViewPage;