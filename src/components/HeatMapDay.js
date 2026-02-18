import React from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { getHeatMapColor } from '../utils/colors';

export const HeatMapDay = ({ date, rating, onPress }) => {
  const backgroundColor = getHeatMapColor(rating || 0);
  const day = date.day;

  return (
    <View
      style={[
        styles.dayContainer,
        { backgroundColor },
      ]}
      onTouchEnd={() => onPress && onPress(date)}
    >
      <Text style={styles.dayText}>{day}</Text>
      {rating > 0 && (
        <Text style={styles.ratingText}>{rating}</Text>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  dayContainer: {
    width: 40,
    height: 40,
    borderRadius: 4,
    alignItems: 'center',
    justifyContent: 'center',
    margin: 2,
  },
  dayText: {
    fontSize: 14,
    color: '#333',
  },
  ratingText: {
    fontSize: 10,
    color: '#fff',
    fontWeight: 'bold',
  },
});
