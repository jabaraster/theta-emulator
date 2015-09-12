'use strict';

const React = require('react/addons');

const Bootstrap = require('react-bootstrap');
const Button = require('react-bootstrap');
const Glyphicon = Bootstrap.Glyphicon;
const Modal = Bootstrap.Modal;
const CSSTransitionGroup = React.addons.CSSTransitionGroup;
/*
const ButtonToolbar = Bootstrap.ButtonToolbar;
const ButtonGroup = Bootstrap.ButtonGroup;
const Button = Bootstrap.Button;
const Label = Bootstrap.Label;
const Input = Bootstrap.Input;
const TabbedArea = Bootstrap.TabbedArea;
const TabPane = Bootstrap.TabPane;
*/
const ajax = require('./component/app-ajax.js');
const classNames = require('classnames');

const PropertyList = React.createClass({
    propTypes: {
        data: React.PropTypes.array.isRequired,
        onPropertyClick: React.PropTypes.func.isRequired,
    },
    handleClick: function(pProperty) {
        this.props.onPropertyClick({ property: pProperty });
    },
    render: function() {
        return (
            <ul className="properties">
                {this.props.data.map((pProperty, idx) => {
                    return (
                        <li key={idx} className="property" onClick={this.handleClick.bind(this, pProperty)}>{pProperty.Name}</li>
                    );
                })}
            </ul>
        );
    },
});

const RoomList = React.createClass({
    propTypes: {
        data: React.PropTypes.array.isRequired,
        onRoomClick: React.PropTypes.func.isRequired,
        onAddRoomClick: React.PropTypes.func.isRequired,
    },
    handleRoomClick: function(e) {
        this.props.onRoomClick(e);
    },
    handleAddRoomClick: function() {
        this.props.onAddRoomClick();
    },
    render: function() {
        const rows = this.props.data.map((pRoom, idx) => {
                        return (
                            <li key={idx} className="room">
                                <Room data={pRoom} onRoomClick={this.props.onRoomClick} />
                            </li>
                        );
                    });
        const add = (
                <li key={this.props.data.length} className="room" onClick={this.handleAddRoomClick}>
                    <Glyphicon glyph="plus" className="add-room" />
                </li>
        );
        rows.push(add);
        return (
            <ul className="rooms">
                {rows}
            </ul>
        );
    },
});

const Room = React.createClass({
    propTypes: {
        data: React.PropTypes.object.isRequired,
        onRoomClick: React.PropTypes.func.isRequired,
    },
    handleRoomClick: function() {
        this.props.onRoomClick({ room: this.props.data });
    },
    render: function() {
        return (
            <div className="Room" onClick={this.handleRoomClick}>
                <img className="room-thumbnail" src="/pano/pano.jpg" />
                <div className="room-name">{this.props.data.Name}</div>
            </div>
        );
    },
});

const BackButton = React.createClass({
    propTypes: {
        onClick: React.PropTypes.func.isRequired,
    },
    handleClick: function() {
        this.props.onClick();
    },
    render: function() {
        return (
            <CSSTransitionGroup component="div" transitionName="fade" className="back">
                <Glyphicon glyph="circle-arrow-left" onClick={this.handleClick} />
            </CSSTransitionGroup>
        );
    },
});

const RoomEditor = React.createClass({
    render: function() {
        return (
            <div>
                <h1>RoomEditor</h1>
            </div>
        );
    },
});

const Page = React.createClass({
    getInitialState: function() {
        return {
            properties: InitialData,
            currentProperty: null,
            propertyListVisible: true,
            viewMode: 'select-property',
        };
    },
    handlePropertyClick: function(e) {
        this.setState({ currentProperty: e.property, viewMode: 'select-room' });
    },
    handleBackClickOnSelectRoom: function() {
        this.setState({ viewMode: 'select-property' });
    },
    handleRoomClick: function(e) {
        this.setState({ viewMode: 'view-room' });
    },
    handleBackClickOnViewRoom: function() {
        this.setState({ viewMode: 'select-room' });
    },
    handleBackClickOnEditRoom: function() {
        this.setState({ viewMode: 'select-room' });
    },
    handleAddRoomClick: function() {
        this.setState({ viewMode: 'edit-room' });
    },
    render: function() {
        return (
            <div className="Page">
                <CSSTransitionGroup transitionName="fade" component="div" className="PropertyList">
                    {this.state.viewMode === 'select-property' ?
                    <PropertyList data={this.state.properties} onPropertyClick={this.handlePropertyClick} />
                    : null}
                </CSSTransitionGroup>

                <CSSTransitionGroup transitionName="fade" component="div">
                    {this.state.viewMode === 'select-room' ?
                    <BackButton onClick={this.handleBackClickOnSelectRoom} />
                    : null}
                </CSSTransitionGroup>
                <CSSTransitionGroup transitionName="fade" component="div" className="RoomList">
                    {this.state.viewMode === 'select-room' ?
                    <RoomList data={this.state.currentProperty ? this.state.currentProperty.Rooms : []}
                              onRoomClick={this.handleRoomClick}
                              onAddRoomClick={this.handleAddRoomClick}
                    />
                    : null}
                </CSSTransitionGroup>

                <CSSTransitionGroup transitionName="fade" component="div">
                    {this.state.viewMode === 'edit-room' ?
                    <BackButton onClick={this.handleBackClickOnEditRoom} />
                    : null}
                </CSSTransitionGroup>
                <CSSTransitionGroup transitionName="fade" component="div" className="RoomEditor">
                    {this.state.viewMode === 'edit-room' ?
                    <RoomEditor />
                    : null}
                </CSSTransitionGroup>

                <CSSTransitionGroup transitionName="fade" component="div">
                    {this.state.viewMode === 'view-room' ?
                    <BackButton onClick={this.handleBackClickOnViewRoom} />
                    : null}
                </CSSTransitionGroup>
            </div>
        );
    },
});

const InitialData = [
    {   Name: '森都心プラザ',
        Rooms: [{ Name: '２階プラザ' }, { Name: '２階図書館' }, { Name: '３階ビジネス図書館' }, { Name: '５階大ホール' }, ]
    },
    {   Name: '森都心プラザ1',
        Rooms: [{ Name: '２階プラザ' }, { Name: '２階図書館' }, { Name: '３階ビジネス図書館' }, { Name: '５階大ホール' }, ]
    },
    {   Name: '森都心プラザ2',
        Rooms: [{ Name: '２階プラザ' }, { Name: '２階図書館' }, { Name: '３階ビジネス図書館' }, { Name: '５階大ホール' }, ]
    },
    {   Name: '森都心プラザ3',
        Rooms: [{ Name: '２階プラザ' }, { Name: '２階図書館' }, { Name: '３階ビジネス図書館' }, { Name: '５階大ホール' }, ]
    },
    {   Name: 'ザ・熊本タワー',
        Rooms: [{ Name: '101号室' }, { Name: '905号室' }, ]
    },
    {   Name: 'ザ・熊本タワー1',
        Rooms: [{ Name: '101号室' }, { Name: '905号室' }, ]
    },
    {   Name: 'ザ・熊本タワー2',
        Rooms: [{ Name: '101号室' }, { Name: '905号室' }, ]
    },
    {   Name: 'ザ・熊本タワー3',
        Rooms: [{ Name: '101号室' }, { Name: '905号室' }, ]
    },
    ];

React.render(
    <Page />,
    document.getElementById('page')
);
