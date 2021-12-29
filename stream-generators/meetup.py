import json
import requests
import boto3
import decorators


@decorators.firehose(
    stream_name="MEETUP-RSVP-2-REDSHIFT",
    client=boto3.client("firehose", region_name="eu-central-1"),
)
@decorators.console
def publish(data_point):
    data = json.loads(data_point)

    filtered_data = {
        "rsvp_id": data["rsvp_id"],
        "ts": data["mtime"],
        "event_time": data["event"]["time"],
        "name": data["member"]["member_name"],
        "response": data["response"],
        "group_city": data["group"]["group_city"],
        "group_country": data["group"]["group_country"],
        "group_name": data["group"]["group_name"],
        "venue_name": data.get("venue", {"venue_name": ""})["venue_name"],
    }
    return filtered_data


def generate():

    r = requests.get("https://stream.meetup.com/2/rsvps", stream=True)

    if r.encoding is None:
        r.encoding = "utf-8"

    for line in r.iter_lines(decode_unicode=True):
        publish(line)


if __name__ == "__main__":
    generate()
