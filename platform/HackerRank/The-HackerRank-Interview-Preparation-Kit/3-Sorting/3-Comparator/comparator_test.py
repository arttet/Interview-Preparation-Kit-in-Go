import sys
import tempfile
import unittest

import comparator


class TestComparator(unittest.TestCase):
    def test_ok(self):
        tests = [
            ("input/input00.txt", "output/output00.txt"),
            ("input/input06.txt", "output/output06.txt"),
            ("input/input07.txt", "output/output07.txt"),
        ]

        for index, test in enumerate(tests):
            with open(test[0], 'r') as stdin, tempfile.NamedTemporaryFile('r+') as stdout, open(test[1], 'r') as f:
                expected = f.read().rstrip("\n")
                test_case = stdin.read().rstrip("\n")
                stdin.seek(0)

                sys.stdin = stdin
                sys.stdout = stdout

                comparator.main()

                stdout.seek(0)
                actual = stdout.read().rstrip("\n")
                self.assertEqual(expected, actual,
                                 f'Test Case #{index}: {test_case}')


if __name__ == '__main__':
    unittest.main()
