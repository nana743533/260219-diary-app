import React, { useState, useEffect } from 'react';
import { View, Text, StyleSheet, Alert, TouchableOpacity } from 'react-native';
import { useFocusEffect } from '@react-navigation/native';
import { Calendar } from 'react-native-calendars';
import { storageService } from '../services/storageService';
import { generateMarkedDates, getHeatMapColor } from '../utils/colors';

export default function CalendarScreen({ navigation }) {
  const [diaries, setDiaries] = useState([]);
  const [markedDates, setMarkedDates] = useState({});
  const [selectedDate, setSelectedDate] = useState(null);

  const loadDiaries = async () => {
    const allDiaries = await storageService.getAllDiaries();
    setDiaries(allDiaries);
    setMarkedDates(generateMarkedDates(allDiaries));
  };

  useEffect(() => {
    loadDiaries();
  }, []);

  useFocusEffect(
    React.useCallback(() => {
      loadDiaries();
    }, [])
  );

  const loadDiaries = async () => {
    const allDiaries = await storageService.getAllDiaries();
    setDiaries(allDiaries);
    setMarkedDates(generateMarkedDates(allDiaries));
  };

  const handleDayPress = (day) => {
    const diary = diaries.find((d) => d.date === day.dateString);
    setSelectedDate(day.dateString);

    if (diary) {
      Alert.alert(
        `${day.dateString}の日記`,
        `総合評価: ${diary.rating}\n進捗: ${diary.progress}\n起床: ${diary.wakeUpTime}\n睡眠: ${diary.sleepTime}\nメモ: ${diary.memo || 'なし'}`,
        [
          { text: '閉じる', style: 'cancel' },
          { text: '編集', onPress: () => navigation.navigate('DiaryEntry', { date: day.dateString }) },
        ]
      );
    } else {
      Alert.alert(
        `${day.dateString}`,
        'まだ日記がありません。記録しますか？',
        [
          { text: 'キャンセル', style: 'cancel' },
          { text: '記録する', onPress: () => navigation.navigate('DiaryEntry', { date: day.dateString }) },
        ]
      );
    }
  };

  const getDayComponent = () => {
    return ({ date, state }) => {
      const diary = diaries.find((d) => d.date === date.dateString);
      const rating = diary ? diary.rating : 0;
      const backgroundColor = getHeatMapColor(rating);

      return (
        <TouchableOpacity
          style={[
            styles.dayContainer,
            state === 'disabled' && styles.disabledDay,
            { backgroundColor },
          ]}
          onPress={() => handleDayPress(date)}
        >
          <Text style={styles.dayText}>{date.day}</Text>
          {rating > 0 && <Text style={styles.ratingText}>{rating}</Text>}
        </TouchableOpacity>
      );
    };
  };

  return (
    <View style={styles.container}>
      <View style={styles.header}>
        <Text style={styles.headerTitle}>日記カレンダー</Text>
        <Text style={styles.headerSubtitle}>評価が緑色の濃さで表示されます</Text>
      </View>

      <Calendar
        current={new Date().toISOString().split('T')[0]}
        dayComponent={getDayComponent()}
        markingType="custom"
        theme={{
          backgroundColor: '#ffffff',
          calendarBackground: '#ffffff',
          textSectionTitleColor: '#b6c1cd',
          selectedDayBackgroundColor: '#2196F3',
          selectedDayTextColor: '#ffffff',
          todayTextColor: '#00adf5',
          dayTextColor: '#2d3436',
          textDisabledColor: '#d1d2d6',
          arrowColor: '#2196F3',
          monthTextColor: '#333',
          textDayFontWeight: '300',
          textMonthFontWeight: 'bold',
          textDayHeaderFontWeight: '300',
        }}
      />

      <View style={styles.legend}>
        <Text style={styles.legendTitle}>凡例</Text>
        <View style={styles.legendRow}>
          <View style={[styles.legendBox, { backgroundColor: getHeatMapColor(0) }]} />
          <Text style={styles.legendText}>なし</Text>
          <View style={[styles.legendBox, { backgroundColor: getHeatMapColor(1) }]} />
          <Text style={styles.legendText}>1</Text>
          <View style={[styles.legendBox, { backgroundColor: getHeatMapColor(2) }]} />
          <Text style={styles.legendText}>2</Text>
          <View style={[styles.legendBox, { backgroundColor: getHeatMapColor(3) }]} />
          <Text style={styles.legendText}>3</Text>
          <View style={[styles.legendBox, { backgroundColor: getHeatMapColor(4) }]} />
          <Text style={styles.legendText}>4</Text>
          <View style={[styles.legendBox, { backgroundColor: getHeatMapColor(5) }]} />
          <Text style={styles.legendText}>5</Text>
        </View>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    paddingTop: 20,
  },
  header: {
    padding: 20,
    alignItems: 'center',
  },
  headerTitle: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#333',
  },
  headerSubtitle: {
    fontSize: 14,
    color: '#666',
    marginTop: 5,
  },
  dayContainer: {
    width: 36,
    height: 36,
    borderRadius: 4,
    alignItems: 'center',
    justifyContent: 'center',
    margin: 1,
  },
  disabledDay: {
    opacity: 0.3,
  },
  dayText: {
    fontSize: 14,
    color: '#333',
  },
  ratingText: {
    fontSize: 10,
    color: '#fff',
    fontWeight: 'bold',
    position: 'absolute',
    bottom: 2,
    right: 2,
  },
  legend: {
    padding: 20,
    borderTopWidth: 1,
    borderTopColor: '#eee',
  },
  legendTitle: {
    fontSize: 14,
    fontWeight: '600',
    color: '#555',
    marginBottom: 10,
  },
  legendRow: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-around',
  },
  legendBox: {
    width: 24,
    height: 24,
    borderRadius: 4,
  },
  legendText: {
    fontSize: 12,
    color: '#666',
    marginLeft: 4,
  },
});
