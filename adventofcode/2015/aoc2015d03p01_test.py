import unittest
from aoc2015d03p01 import execute

class TestAoc201503P01(unittest.TestCase):

    def test_example001(self):
        res = execute('>')
        self.assertEqual(2, res)

    def test_example002(self):
        res = execute('^>v<')
        self.assertEqual(4, res)

    def test_for_real(self):
        input = ''
        with open('input-aoc2015d03p01.txt', 'r') as f:
            input = f.read()
        res = execute(input)
        self.assertEqual(2565, res)
    
if __name__ == '__main__':
    unittest.main()