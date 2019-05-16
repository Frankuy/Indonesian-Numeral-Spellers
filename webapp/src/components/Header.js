import React, { Component } from 'react'

class Header extends Component {
    render() {
        return (
            <nav className="navbar navbar-dark myNav">
                <div className="container-fluid">
                    <div className="navbar-header">
                        <a className="judul navbar-brand" href="https://github.com/Frankuy/Indonesian-Numeral-Spellers">Indonesian Numeral Speller</a>
                    </div>
                </div>
            </nav>
        );
    }
}
        
export default Header