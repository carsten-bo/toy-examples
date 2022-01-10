import functools
import json


def console(func):
    @functools.wraps(func)
    def decorator(*args, **kwargs):
        retval = func(*args, **kwargs)
        print(retval)
        return retval

    return decorator


def kinesis_firehose(stream_name, client):
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


def kinesis_data_stream(stream_name, client, **outer_kwargs):
    def wrapped_decorator(func):
        def decorator(*args, **kwargs):
            retval = func(*args, **kwargs)
            client.put_record(
                StreamName=stream_name,
                Data=json.dumps(retval).encode("UTF-8"),
                **outer_kwargs
            )
            return retval

        return decorator

    return wrapped_decorator

def kafka(stream_name, client, **outer_kwargs):
    def wrapped_decorator(func):
        def decorator(*args, **kwargs):
            retval = func(*args, **kwargs)
            client.put_record(
                StreamName=stream_name,
                Data=json.dumps(retval).encode("UTF-8"),
                **outer_kwargs
            )
            return retval

        return decorator

    return wrapped_decorator

