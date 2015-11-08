/// <reference path='all.d.ts' />
class Group extends React.Component<Model.Group, any> {
  constructor() {
    super();
  }

  render() {
    const tiles = this.props.tiles.map(tile => {
      return <Tile title={tile.title} description={tile.description} values={tile.values} key={tile.title}/>
    });

    return (
      <div className="row card blue-grey darken-3">
        <div className="card-content">
          <div className="card-title card-title-group grey-text">{this.props.name}</div>
          {tiles}
          </div>
        </div>
    )
  }
}

class Tile extends React.Component<Model.Tile, any> {
  constructor() {
    super()
  }

  private createTile(value: Model.ValueWithTimestamp, index: number): JSX.Element {
    const ago = moment.unix(value.timestamp).fromNow();
    const key = 'dot-' + index;
    var className = 'dot tooltipped new-item'
    if (value.value === 0) className += " hide";
    if (value.value === 1) className += " green accent-3";
    if (value.value === 2) className += " red pulse accent-1";
    return <li key={ key }><span className={className} title={ago} > dot </span></li>
  }

  render() {
    const dots = this.props.values
      .sort((a, b) => b.timestamp - a.timestamp)
      .map(this.createTile)
    return (
      <div className="col s12 m6 l4">
        <div className="card teal darken-2">
          <div className="card-content white-text">
            <ol>
              {dots}
            </ol>
            <span className="card-title">{this.props.title}</span>
            <p>{this.props.description}</p>
            </div>
          </div>
        </div>
    )
  }
}
