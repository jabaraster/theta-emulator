'use strict';

const React = require('react');

const Bootstrap = require('react-bootstrap');
const Button = Bootstrap.Button;
const Glyphicon = Bootstrap.Glyphicon;
const ajax = require('superagent');
const classNames = require('classnames');

const Page = React.createClass({
    getInitialState: function() {
        return {
            url: '',
            response: '',
        };
    },
    handleUrlChange: function() {
        this.setState({ url: React.findDOMNode(this.refs.url).value });
    },
    handleClick: function() {
        const url = React.findDOMNode(this.refs.url).value;
        const json = React.findDOMNode(this.refs.request).value;
        if (!url) {
            return;
        }
        ajax.post(url).
            send(json).
            set('content-type', 'application/json').
            end((err, res) => {
                if (err) {
                    console.log(err);
                    console.log(res);
                    alert('エラー発生. ブラウザのコンソールを確認すること.');
                    return;
                }
                if (!res.ok) {
                    console.log(res);
                    alert('ブラウザのコンソールを確認すること.');
                    return;
                }
                this.setState({ response: res.text });
            });
    },
    render: function() {
        return (
            <div className="Page container">
                <div className="form-group">
                    <label htmlFor="url-text">URL:</label>
                    <input type="text" id="url-text" className="form-control" ref="url"
                           value={this.state.url}
                           onChange={this.handleUrlChange}
                    />
                </div>
                <div className="form-group">
                    <label htmlFor="request-text">Request:</label>
                    <textarea id="request-text" className="form-control" ref="request"></textarea>
                </div>
                <div className="form-group">
                    <label htmlFor="response-text">Response:</label>
                    <textarea id="response-text" className="form-control" ref="response" value={this.state.response}></textarea>
                </div>
                <Button bsStyle="primary" onClick={this.handleClick} disabled={this.state.url.length === 0}>
                    <Glyphicon glyph="ok"/>
                </Button>
            </div>
        );
    },
});

React.render(
    <Page />,
    document.getElementById('page')
);
