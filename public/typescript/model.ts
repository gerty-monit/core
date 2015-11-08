declare module Model {
  interface ReactItem {
    key?: string;
  }

  interface Dashboard {
    groups: Array<Group>;
  }

  interface Group extends ReactItem {
    name: string;
    tiles: Array<Tile>;
  }

  interface Tile extends ReactItem {
    title: string;
    description: string;
    values: Array<ValueWithTimestamp>;
  }

  interface ValueWithTimestamp {
    value: number;
    timestamp: number;
  }
}