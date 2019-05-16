import React, { Component } from 'react'
import Header from './components/Header'
import MainContent from './components/MainContent'
import Footer from './components/Footer'

class MyApp extends Component {
    render() {
        return (
            <div>
                <center>
                <Header />
                <MainContent />
                <Footer />
                </center>
            </div>
        );
    }
}

export default MyApp