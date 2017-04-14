import React from 'react';
import { connect } from 'react-redux';
import { Content } from '../components';
import { getMediaList } from '../helpers/mediaListReq';

class ListPage extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      list: null,
    };
  }

  componentWillMount() {
    getMediaList(this);
  }

  subtitleComponent = (media) => {
    if (media.subtitle != "") {
      return (
        <span className="badge"><i className="fa fa-fw fa-language" /></span>
      );
    } else {
      return (
        <span />
      );
    }
  }

  render() {
    // styles
    const col1Style = {width: "100px"};
    const col2Style = {width: "100px"};


    // TODO: filter by mimetype
    const mediaList = this.state.list ? this.state.list.Data.map(media =>
      <a href={`/media/${media.path}`} className="list-group-item" key={media.name}>
        <h4 class="list-group-item-heading">{media.name} {this.subtitleComponent(media)}</h4>
        
        <p class="list-group-item-text">{media.mimetype}</p>
      </a>,
    ) : null;
    return (
        <div className="col-md-12">
          <h1>Media List</h1>
          <div className="list-group">
            {mediaList}
          </div>
        </div>
    );
  }
}

const mapStateToProps = state => ({
  user: state.user,
});

const ConnectedListPage = connect(mapStateToProps)(ListPage);
export default ConnectedListPage;
