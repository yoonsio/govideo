import React from 'react';
import Sidebar from './Sidebar';

const Content = ({ children }) => (
  <div>
    <Sidebar />
    <div className="main container-fluid">
      {children}
    </div>
  </div>
);

Content.propTypes = {
  children: React.PropTypes.element.isRequired,
};

export default Content;
