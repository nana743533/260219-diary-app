// GitHub風の緑色グラデーション（評価1-5）
export const getHeatMapColor = (rating) => {
  const colors = {
    0: '#ebedf0', // データなし（グレー）
    1: '#c6e48b', // 評価1（薄い緑）
    2: '#7bc96f', // 評価2
    3: '#239a3b', // 評価3
    4: '#196127', // 評価4
    5: '#0e4429', // 評価5（最も濃い緑）
  };
  return colors[rating] || colors[0];
};

// 評価値に基づくマーカーデータを生成
export const generateMarkedDates = (diaries) => {
  const markedDates = {};
  diaries.forEach((diary) => {
    const color = getHeatMapColor(diary.rating);
    markedDates[diary.date] = {
      marked: true,
      selectedColor: color,
      selectedTextColor: diary.rating > 0 ? '#fff' : '#666',
    };
  });
  return markedDates;
};
