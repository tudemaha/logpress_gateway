export default interface SensorData {
  ts: string;
  device: string;
  co: number;
  humidity: number;
  lpg: number;
  temp: number;
  smoke: number;
  motion: boolean;
  light: boolean;
}
