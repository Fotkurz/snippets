import unittest
from aoc2015d04p01 import execute

class TestAoc201504P01(unittest.TestCase):

    def test_example001(self):
        res = execute('abcdef')
        self.assertEqual(609043, res)

    def test_example002(self):
        res = execute('pqrstuv')
        self.assertEqual(1048970, res)

    def test_for_real(self):
        res = execute('bgvyzdsv')
        self.assertEqual(254575, res)
    
if __name__ == '__main__':
    unittest.main()