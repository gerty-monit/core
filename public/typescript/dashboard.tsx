/// <reference path='all.d.ts' />
/// <reference path='cards.tsx' />

class Dashboard extends React.Component<{ interval: number }, Model.Dashboard> {
  constructor() {
    super()
    this.state = { groups: [] };
  }

  componentDidMount() {
    this.fetchAndUpdate();
    setInterval(this.fetchAndUpdate.bind(this), this.props.interval)
  }

  fetchAndUpdate() {
    $.get('/api/v1/monitors', (data: Array<Model.Group>) => {
      this.setState({ groups: data });
    })
  }

  render() {
    const groups = this.state.groups.map(it => <Group name={it.name} tiles={it.tiles} key={it.name}/>)
    return (
      <div className="container">
        <div className="section">
          {groups}
        </div>
      </div>
    )
  }
}

ReactDOM.render(
  <Dashboard interval= { 5000 }/>,
  document.getElementById('content')
);