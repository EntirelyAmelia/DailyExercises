import random
import sys

letters = {
	'c': 'bcdfghjklmnpqrstvwxyz',
	'v': 'aeiou'
}

def validate(pattern):
	for char in pattern:
		if char not in ['c', 'v', 'C', 'V']:
			return False
	return True

def replace(char):
	new_char = random.choice(letters.get(char.lower(), char))
	return new_char if char.islower() else new_char.upper()
	
if __name__ == '__main__':
	pattern = sys.argv[1]
	if validate(pattern):
		print(''.join(replace(char) for char in pattern))
	else:
		print('Invalid input. Pattern must consist only of \'c\', \'v\', \'C\', or \'V\'.')
