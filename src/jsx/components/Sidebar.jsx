import React from 'react';
import { Link } from 'react-router';

export default class Sidebar extends React.Component {
    render() {
        return (
            <nav className="sidebar navbar navbar-default">
                <div className="sidebar-container container-fluid">
                    <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#main-navbar-collapse" aria-expanded="false">
                        <span className="sr-only">Toggle navigation</span>
                        <span className="icon-bar"></span>
                        <span className="icon-bar"></span>
                        <span className="icon-bar"></span>
                    </button>
                    <div className="sidebar-title-mobile visible-xs">
                        <h4><i className="fa fa-fw fa-video-camera"></i>GoVideo</h4>
                    </div>
                </div>
                <div className="sidebar-nav collapse navbar-collapse" id="main-navbar-collapse">
                    <ul className="nav navbar-nav">
                        <li>
                            <Link to="/"><i className="fa fa-fw fa-home"></i> Home</Link>
                            <Link to="/login"><i className="fa fa-fw fa-user"></i> Login</Link>
                        </li>
                    </ul>
                </div>
            </nav>
        )
    }
}
