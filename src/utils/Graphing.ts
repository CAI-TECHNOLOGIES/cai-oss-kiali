import { TimeSeries } from '../types/Metrics';

export default {
  toC3Columns(matrix?: TimeSeries[], title?: string) {
    if (!matrix || matrix.length === 0) {
      return [['x'], [title || '']];
    }

    // xseries are timestamps. Timestamps are taken from the first series and assumed
    // that all series have the same timestamps.
    let xseries: any = ['x'];
    xseries = xseries.concat(matrix[0].values.map(dp => dp[0] * 1000));

    // yseries are the values of each serie.
    let yseries: any[] = matrix.map(mat => {
      let serie: any = [title || mat.name];
      return serie.concat(mat.values.map(dp => dp[1]));
    });

    // timestamps + data is the format required by C3 (all concatenated: an array with arrays)
    return [xseries, ...yseries];
  },

  toC3ValueColumns(matrix: TimeSeries[], title?: string) {
    return matrix.map(mat => {
      let yseries: any = [title || mat.name];
      return yseries.concat(mat.values.map(dp => dp[1]));
    });
  }
};
