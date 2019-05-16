import React, { Component } from 'react';
import ReactDOM from 'react-dom';

function numberWithCommas(x) {
    if (!isNaN(x)) return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

class MainContent extends Component {
    state = {
        value: '',
        result: []
    }

    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
        this.getSpell = this.getSpell.bind(this);
        this.getNumeral = this.getNumeral.bind(this);
        this.playVoice = this.playVoice.bind(this);
    }

    playVoice(event) {
        var msg;
        if (this.state.result.status === "OK") {
            if (this.state.result.text !== undefined) {
                msg = new SpeechSynthesisUtterance(this.state.result.text);
            }
            if (this.state.result.number !== undefined) {
                msg = new SpeechSynthesisUtterance(this.state.result.number);
            }
        }

        msg.lang = 'id';
        msg.rate = 0.8;
        window.speechSynthesis.speak(msg);

        event.preventDefault();
    }

    handleChange(event) {
        this.setState({ value: event.target.value });
    }

    getSpell(event) {
        fetch("http://localhost:8080/spell?number=" + this.state.value)
            .then(response => response.json())
            .then(resJson => this.setState({ result: resJson }))
            .catch(console.log);

        ReactDOM.render(<img onClick={this.playVoice} id='voice' className="" src='https://cdn3.iconfinder.com/data/icons/multimedia/100/volume_up-512.png' alt="voice"/>, document.getElementById("voiceCont"))
        event.preventDefault()
    }

    getNumeral(event) {
        fetch("http://localhost:8080/read", {
            method: 'post',
            body: JSON.stringify({
                text: this.state.value,
            })
        })
            .then(response => response.json())
            .then(resJson => this.setState({ result: resJson }))
            .catch(console.log);

        ReactDOM.render(<img onClick={this.playVoice} id='voice' className="" src='https://cdn3.iconfinder.com/data/icons/multimedia/100/volume_up-512.png' alt="voice"/>, document.getElementById("voiceCont"));
        event.preventDefault()
    }

    render() {
        return (
            <div className="container-fluid myContainer">
                <div className="resultContainer">
                    <div id="res">
                        <span className="result"> {numberWithCommas(parseInt(this.state.result.number,10))} {this.state.result.text}</span>
                        <span id='voiceCont'>
                        </span>
                    </div>
                </div>
                <div className="form-group textForm">
                    <input id="input" className="form-control input-text" placeholder="Enter number" type="text" value={this.state.value} onChange={this.handleChange} />
                </div>
                <div className="buttonCont">
                    <div className="buttonn">   
                        <input className="btn btn-warning btn-lg" type="button" value="Spell" onClick={this.getSpell} />
                    </div>
                    <div className="dividerAxis"></div>
                    <div className="buttonn">
                        <input className="btn btn-warning btn-lg" type="button" value="Numeral" onClick={this.getNumeral} />
                    </div>
                </div>
            </div>
        );
    }
}

export default MainContent