import React from 'react';
import { ActivityIndicator, StyleSheet, Text, TouchableOpacity, TouchableOpacityProps } from 'react-native';
import { ShortcutProps, defaultShortcuts } from '@/styles/shortcuts';

interface ButtonProps extends ShortcutProps, TouchableOpacityProps{
	variant?: 'contained' | 'outlined' | 'ghost';
	isLoading?: boolean;
}

export const Button = ({
	onPress,
	children,
	variant = "contained",
	isLoading,
	...restProps
}: ButtonProps) => {
	return (
		<TouchableOpacity
		disabled={isLoading}
		onPress={onPress}
		style={[
			defaultShortcuts(restProps),
			styles[variant].button,
			isLoading && disabled.button
	
		]}
		{...restProps}
		>
			{isLoading ? 
		<ActivityIndicator animating={isLoading} size={22}/>:
		<Text style={styles[variant].text}>{children}</Text>	
		}
		</TouchableOpacity>
	);
};

