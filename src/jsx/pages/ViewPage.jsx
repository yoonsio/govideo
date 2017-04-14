import React from 'react';
import { connect } from 'react-redux';
import { Content } from '../components';
import { getMedia } from '../helpers/mediaReq';

class ViewPage extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            media: null,
            path: null,
            subtitle: null
        };
    }

    componentWillMount() {
        getMedia(this, this.props.params.path);
    }

    videoComponent = () => {
        if (this.state.media == null) {
            return (
                <h2>Loading media...</h2>
            );    
        }
        let subtitleComponent = null;
        if (this.state.subtitle) {
            subtitleComponent = <track kind="captions" label="English captions" src={this.state.subtitle} srclang="en" default />;
        }
        return (
            <video controls>
                <source src={this.state.path} type={this.state.media.mimetype} />
                {subtitleComponent}
            </video>
        );    
    }

    render() {
        return (
            <div>
                {this.videoComponent()}
            </div>
        )
    }
}

const mapStateToProps = state => ({
  user: state.user,
});

const ConnectedViewPage = connect(mapStateToProps)(ViewPage);
export default ConnectedViewPage;