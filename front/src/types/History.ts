export type History = {
  Worlds: World[];
  Configs: Config;
};

export type Cell = {
  IsLive: boolean;
};

export type World = {
  Cells: Cell[];
};

type Config = {
  Debug: boolean;
  GenCap: number;
  Row: number;
  Col: number;
};
