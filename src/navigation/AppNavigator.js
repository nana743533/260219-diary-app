import React from 'react';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { NavigationContainer } from '@react-navigation/native';
import CalendarScreen from '../screens/CalendarScreen';
import DiaryEntryScreen from '../screens/DiaryEntryScreen';

const Tab = createBottomTabNavigator();

export default function AppNavigator() {
  return (
    <NavigationContainer>
      <Tab.Navigator
        screenOptions={{
          headerStyle: {
            backgroundColor: '#2196F3',
          },
          headerTintColor: '#fff',
          headerTitleStyle: {
            fontWeight: 'bold',
          },
          tabBarStyle: {
            backgroundColor: '#fff',
            borderTopWidth: 1,
            borderTopColor: '#eee',
          },
          tabBarActiveTintColor: '#2196F3',
          tabBarInactiveTintColor: '#999',
        }}
      >
        <Tab.Screen
          name="Calendar"
          component={CalendarScreen}
          options={{ title: 'カレンダー' }}
        />
        <Tab.Screen
          name="DiaryEntry"
          component={DiaryEntryScreen}
          options={{ title: '記録' }}
        />
      </Tab.Navigator>
    </NavigationContainer>
  );
}
