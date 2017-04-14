import React from 'react';
import Navbar from './Navbar';

const Content = ({ children }) => (
  <div>
    <Navbar />
    <div className="main container">
      <div className="row">
        {children}
      </div>
    </div>
  </div>
);

Content.propTypes = {
  children: React.PropTypes.element.isRequired,
};

export default Content;
