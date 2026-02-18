import AsyncStorage from '@react-native-async-storage/async-storage';

const DIARY_KEY_PREFIX = 'diary_';

export const storageService = {
  // 日記を保存
  saveDiary: async (date, diaryData) => {
    try {
      const key = `${DIARY_KEY_PREFIX}${date}`;
      const data = {
        ...diaryData,
        date,
        updatedAt: new Date().toISOString(),
      };
      await AsyncStorage.setItem(key, JSON.stringify(data));
      return { success: true };
    } catch (error) {
      console.error('保存失敗:', error);
      return { success: false, error };
    }
  },

  // 特定の日記を取得
  getDiary: async (date) => {
    try {
      const key = `${DIARY_KEY_PREFIX}${date}`;
      const jsonValue = await AsyncStorage.getItem(key);
      return jsonValue != null ? JSON.parse(jsonValue) : null;
    } catch (error) {
      console.error('読み込み失敗:', error);
      return null;
    }
  },

  // 全日記を取得
  getAllDiaries: async () => {
    try {
      const keys = await AsyncStorage.getAllKeys();
      const diaryKeys = keys.filter((key) => key.startsWith(DIARY_KEY_PREFIX));
      const values = await AsyncStorage.multiGet(diaryKeys);
      return values
        .map(([_, value]) => (value != null ? JSON.parse(value) : null))
        .filter(Boolean);
    } catch (error) {
      console.error('全データ読み込み失敗:', error);
      return [];
    }
  },

  // 期間指定で日記を取得
  getDiariesInRange: async (startDate, endDate) => {
    try {
      const allDiaries = await storageService.getAllDiaries();
      return allDiaries.filter((diary) => {
        return diary.date >= startDate && diary.date <= endDate;
      });
    } catch (error) {
      console.error('期間データ読み込み失敗:', error);
      return [];
    }
  },

  // 日記を削除
  deleteDiary: async (date) => {
    try {
      const key = `${DIARY_KEY_PREFIX}${date}`;
      await AsyncStorage.removeItem(key);
      return { success: true };
    } catch (error) {
      console.error('削除失敗:', error);
      return { success: false, error };
    }
  },
};
