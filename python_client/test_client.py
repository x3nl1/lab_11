import time
from client import calculate

def test_calculate():
    for _ in range(5):
        try:
            res = calculate([1, 2, 3])
            assert res["sum"] == 6
            return
        except Exception:
            time.sleep(1)
    assert False