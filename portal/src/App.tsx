import { Button } from '@/components/ui/button';
import { Account } from '@surl/common';

export default function App() {
	let acc: Account = {
		id: 1,
	};

	console.log(acc);

	return (
		<div>
			<Button>portal</Button>
		</div>
	);
}
