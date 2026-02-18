import React, { useState, useEffect } from 'react';
import {
  StyleSheet,
  Text,
  View,
  TextInput,
  TouchableOpacity,
  ScrollView,
  Alert,
} from 'react-native';
import { useRoute } from '@react-navigation/native';
import { storageService } from '../services/storageService';

export default function DiaryEntryScreen() {
  const route = useRoute();
  const todayDate = route.params?.date || new Date().toISOString().split('T')[0];

  const [rating, setRating] = useState(3);
  const [progress, setProgress] = useState('B');
  const [wakeUpTime, setWakeUpTime] = useState('07:00');
  const [sleepTime, setSleepTime] = useState('23:00');
  const [memo, setMemo] = useState('');

  // 既存データがあれば読み込む
  useEffect(() => {
    loadExistingDiary();
  }, [todayDate]);

  const loadExistingDiary = async () => {
    const existing = await storageService.getDiary(todayDate);
    if (existing) {
      setRating(existing.rating);
      setProgress(existing.progress);
      setWakeUpTime(existing.wakeUpTime);
      setSleepTime(existing.sleepTime);
      setMemo(existing.memo);
    }
  };

  const handleSave = async () => {
    const diaryData = {
      date: todayDate,
      rating,
      progress,
      wakeUpTime,
      sleepTime,
      memo,
    };

    const result = await storageService.saveDiary(todayDate, diaryData);

    if (result.success) {
      Alert.alert('保存完了', `${todayDate}の日記を保存しました！`, [
        { text: 'OK', onPress: () => {} }
      ]);
    } else {
      Alert.alert('保存失敗', 'もう一度お試しください');
    }
  };

  return (
    <View style={styles.container}>
      <ScrollView style={styles.scrollView} contentContainerStyle={styles.scrollContent}>
        <Text style={styles.title}>{todayDate} 日記記録</Text>

        {/* 総合評価 */}
        <View style={styles.section}>
          <Text style={styles.label}>総合評価 (1-5)</Text>
          <View style={styles.ratingContainer}>
            {[1, 2, 3, 4, 5].map((num) => (
              <TouchableOpacity
                key={num}
                style={[
                  styles.ratingButton,
                  rating === num && styles.ratingButtonActive,
                ]}
                onPress={() => setRating(num)}
              >
                <Text
                  style={[
                    styles.ratingText,
                    rating === num && styles.ratingTextActive,
                  ]}
                >
                  {num}
                </Text>
              </TouchableOpacity>
            ))}
          </View>
        </View>

        {/* 進捗 */}
        <View style={styles.section}>
          <Text style={styles.label}>進捗</Text>
          <View style={styles.progressContainer}>
            {['A', 'B', 'C'].map((p) => (
              <TouchableOpacity
                key={p}
                style={[
                  styles.progressButton,
                  progress === p && styles.progressButtonActive,
                ]}
                onPress={() => setProgress(p)}
              >
                <Text
                  style={[
                    styles.progressText,
                    progress === p && styles.progressTextActive,
                  ]}
                >
                  {p}
                </Text>
              </TouchableOpacity>
            ))}
          </View>
        </View>

        {/* 起床時間 */}
        <View style={styles.section}>
          <Text style={styles.label}>起床時間</Text>
          <TextInput
            style={styles.input}
            value={wakeUpTime}
            onChangeText={setWakeUpTime}
            placeholder="07:00"
            placeholderTextColor="#999"
          />
        </View>

        {/* 睡眠時間 */}
        <View style={styles.section}>
          <Text style={styles.label}>睡眠時間</Text>
          <TextInput
            style={styles.input}
            value={sleepTime}
            onChangeText={setSleepTime}
            placeholder="23:00"
            placeholderTextColor="#999"
          />
        </View>

        {/* メモ */}
        <View style={styles.section}>
          <Text style={styles.label}>メモ</Text>
          <TextInput
            style={[styles.input, styles.textArea]}
            value={memo}
            onChangeText={setMemo}
            placeholder="今日の出来事などを記録..."
            placeholderTextColor="#999"
            multiline
            numberOfLines={4}
            textAlignVertical="top"
          />
        </View>

        {/* 保存ボタン */}
        <TouchableOpacity style={styles.saveButton} onPress={handleSave}>
          <Text style={styles.saveButtonText}>保存</Text>
        </TouchableOpacity>
      </ScrollView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f5f5f5',
  },
  scrollView: {
    flex: 1,
  },
  scrollContent: {
    padding: 20,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    textAlign: 'center',
    marginVertical: 20,
    color: '#333',
  },
  section: {
    marginBottom: 20,
  },
  label: {
    fontSize: 16,
    fontWeight: '600',
    marginBottom: 10,
    color: '#555',
  },
  ratingContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  ratingButton: {
    width: 60,
    height: 60,
    borderRadius: 30,
    backgroundColor: '#e0e0e0',
    alignItems: 'center',
    justifyContent: 'center',
  },
  ratingButtonActive: {
    backgroundColor: '#FFD700',
  },
  ratingText: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#666',
  },
  ratingTextActive: {
    color: '#fff',
  },
  progressContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  progressButton: {
    width: 80,
    height: 50,
    borderRadius: 8,
    backgroundColor: '#e0e0e0',
    alignItems: 'center',
    justifyContent: 'center',
  },
  progressButtonActive: {
    backgroundColor: '#4CAF50',
  },
  progressText: {
    fontSize: 20,
    fontWeight: 'bold',
    color: '#666',
  },
  progressTextActive: {
    color: '#fff',
  },
  input: {
    backgroundColor: '#fff',
    borderRadius: 8,
    padding: 15,
    fontSize: 16,
    borderWidth: 1,
    borderColor: '#ddd',
  },
  textArea: {
    height: 120,
  },
  saveButton: {
    backgroundColor: '#2196F3',
    borderRadius: 8,
    padding: 18,
    alignItems: 'center',
    marginTop: 10,
    marginBottom: 30,
  },
  saveButtonText: {
    color: '#fff',
    fontSize: 18,
    fontWeight: 'bold',
  },
});
