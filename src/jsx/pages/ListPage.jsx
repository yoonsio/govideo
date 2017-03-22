import React from 'react';
import { connect } from 'react-redux';
import { Content } from '../components';
import { getMediaList } from '../helpers/mediaReq';

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

  render() {
    // TODO: filter by mimetype
    const mediaList = this.state.list ? this.state.list.Data.map(media =>
      <li key={media.name}>{media.name} {media.mimetype}</li>,
    ) : null;
    return (
      <Content>
        <div>
          <h1>Media List</h1>
          <ul>
            {mediaList}
          </ul>
        </div>
      </Content>
    );
  }
}

const mapStateToProps = state => ({
  user: state.user,
});

const ConnectedListPage = connect(mapStateToProps)(ListPage);
export default ConnectedListPage;
