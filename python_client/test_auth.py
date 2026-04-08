import time
from client import get_token, protected

def test_protected():
    for _ in range(5):
        try:
            token = get_token()
            status = protected(token)
            assert status == 200
            return
        except Exception:
            time.sleep(1)
    assert False