import unittest
from aoc2015d04p02 import execute

class TestAoc201504P02(unittest.TestCase):

    def test_for_real(self):
        res = execute('bgvyzdsv')
        self.assertEqual(1038736, res)
    
if __name__ == '__main__':
    unittest.main()