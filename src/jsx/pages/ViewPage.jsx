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
            subtitle_path: null,
        }
    }

    componentWillMount() {
        getMedia(this, this.props.params.path);
    }

    translateType = () => {
        switch (this.state.media.mimetype) {
            case "video/x-matroska":
                return "video/webm";
            default:
                return type;
        }
    }

    subtitleComponent = () => {
        if (this.state.subtitle_path) {
            return (
                <track kind="captions" label="English captions" src="/path/to/captions.vtt" srclang="en" default />
            );
        }
        return (<div />);
    }
    
    render() {
        let fixedWidth = {
            width: "600px"
        };
        let videoComponent;
        if (this.state.media != null) {
            videoComponent = <video controls><source src={this.state.path} type={this.state.media.mimetype} />{subtitleComponent()}</video>
        } else {
            videoComponent = <h2>Loading media...</h2>
        }
        return (
            <div style={fixedWidth}>
                {this.state.media}
            </div>
        )
    }
}

const mapStateToProps = state => ({
  user: state.user,
});

const ConnectedViewPage = connect(mapStateToProps)(ViewPage);
export default ConnectedViewPage;