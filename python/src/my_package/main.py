"""Main module for my_package."""


def my_function_1() -> str:
    """Return a greeting message."""
    return "Hello, World!"


def my_function_2() -> int:
    """Return a constant integer value."""
    return 42


def run() -> None:
    """Validate expected outputs from module functions."""
    if my_function_1() != "Hello, World!":
        err_msg = f"Expected 'Hello, World!', got '{my_function_1()}'"
        raise ValueError(err_msg)

    if my_function_2() != 42:  # noqa: PLR2004
        err_msg = f"Expected 42, got '{my_function_2()}'"
        raise ValueError(err_msg)


if __name__ == "__main__":
    run()
