import unittest
from aoc2015d03p02 import execute

class TestAoc201503P02(unittest.TestCase):

    def test_example001(self):
        res = execute('^v')
        self.assertEqual(3, res)

    def test_example002(self):
        res = execute('^>v<')
        self.assertEqual(3, res)
    
    def test_example003(self):
        res = execute('^v^v^v^v^v')
        self.assertEqual(11, res)

    def test_for_real(self):
        input = ''
        with open('input-aoc2015d03p01.txt', 'r') as f:
            input = f.read()

        res = execute(input)
        self.assertEqual(2639, res)
    
if __name__ == '__main__':
    unittest.main()