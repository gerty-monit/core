var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
/// <reference path='all.d.ts' />
var Group = (function (_super) {
    __extends(Group, _super);
    function Group() {
        _super.call(this);
    }
    Group.prototype.render = function () {
        var tiles = this.props.tiles.map(function (tile) {
            return React.createElement(Tile, {title: tile.title, description: tile.description, values: tile.values, key: tile.title});
        });
        return (React.createElement("div", {className: "row card blue-grey darken-3"}, React.createElement("div", {className: "card-content"}, React.createElement("div", {className: "card-title card-title-group grey-text"}, this.props.name), tiles)));
    };
    return Group;
}(React.Component));
var Tile = (function (_super) {
    __extends(Tile, _super);
    function Tile() {
        _super.call(this);
    }
    Tile.prototype.createTile = function (value, index) {
        var ago = moment.unix(value.timestamp).fromNow();
        var key = 'dot-' + index;
        var className = 'dot tooltipped new-item';
        if (value.value === 0)
            className += " green accent-3";
        if (value.value === 1)
            className += " red pulse accent-1";
        return React.createElement("li", {key: key}, React.createElement("span", {className: className, title: ago}, " dot "));
    };
    Tile.prototype.render = function () {
        var dots = this.props.values
            .sort(function (a, b) { return b.timestamp - a.timestamp; })
            .map(this.createTile);
        var failed = this.props.values.filter(function (it) { return it.value === 1; }).length;
        var allFailed = (failed === this.props.values.length);
        var tileClass = (allFailed) ? "red darken-2 card" : "teal darken-2 card";
        return (React.createElement("div", {className: "col s12 m6 l4"}, React.createElement("div", {className: tileClass}, React.createElement("div", {className: "card-content white-text"}, React.createElement("ol", null, dots), React.createElement("span", {className: "card-title"}, this.props.title), React.createElement("p", null, this.props.description)))));
    };
    return Tile;
}(React.Component));
