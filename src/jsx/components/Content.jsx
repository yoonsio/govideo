import React from 'react';
import Sidebar from './Sidebar';

export default class Content extends React.Component {
    render() {
        return (
            <div>
                <Sidebar/>
                <div className="main container-fluid">
                    {this.props.children}
                </div>
            </div>
        )
    }
}
