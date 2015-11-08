/// <reference path='all.d.ts' />
/// <reference path='cards.tsx' />
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
var Dashboard = (function (_super) {
    __extends(Dashboard, _super);
    function Dashboard() {
        _super.call(this);
        this.state = { groups: [] };
    }
    Dashboard.prototype.componentDidMount = function () {
        this.fetchAndUpdate();
        setInterval(this.fetchAndUpdate.bind(this), this.props.interval);
    };
    Dashboard.prototype.fetchAndUpdate = function () {
        var _this = this;
        $.get('/api/v1/monitors', function (data) {
            _this.setState({ groups: data });
        });
    };
    Dashboard.prototype.render = function () {
        var groups = this.state.groups.map(function (it) { return React.createElement(Group, {"name": it.name, "tiles": it.tiles, "key": it.name}); });
        return (React.createElement("div", {"className": "container"}, React.createElement("div", {"className": "section"}, groups)));
    };
    return Dashboard;
})(React.Component);
ReactDOM.render(React.createElement(Dashboard, {"interval": 5000}), document.getElementById('content'));
