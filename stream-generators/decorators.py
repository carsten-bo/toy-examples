import functools
import json


def console(func):
    @functools.wraps(func)
    def decorator(*args, **kwargs):
        retval = func(*args, **kwargs)
        print(retval)
        return retval

    return decorator


def firehose(stream_name, client):
    def wrapped_decorator(func):
        def decorator(*args, **kwargs):
            retval = func(*args, **kwargs)
            client.put_record(
                DeliveryStreamName=stream_name,
                Record={"Data": json.dumps(retval).encode("UTF-8")},
            )
            return retval

        return decorator

    return wrapped_decorator
