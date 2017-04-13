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

  mediaEntry = (media) => {
    if (media.subtitle != "") {
      return (
        <a href={`/media/${media.path}`}>
          {media.name}
          <i className="fa fa-fw fa-language" />
        </a>
      );
    } else {
      return (
        <a href={`/media/${media.path}`}>
          {media.name} 
        </a>
      );
    }
  }

  render() {
    // TODO: filter by mimetype
    const mediaList = this.state.list ? this.state.list.Data.map(media =>
      <li key={media.name}>
        {this.mediaEntry(media)}
      </li>,
    ) : null;
    return (
        <div>
          <h1>Media List</h1>
          <ul>
            {mediaList}
          </ul>
        </div>
    );
  }
}

const mapStateToProps = state => ({
  user: state.user,
});

const ConnectedListPage = connect(mapStateToProps)(ListPage);
export default ConnectedListPage;
